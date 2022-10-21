package admindb

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
	configuration "xqledger/apirouter/configuration"
	utils "xqledger/apirouter/utils"
	_ "github.com/go-sql-driver/mysql"
)

const componentMessage = "Admin Db Client"
var config = configuration.GlobalConfiguration
var mainConn *sql.DB

func dbConn()  error {
    dbDriver := "mysql"
    dbUser := config.Admindb.Username
    dbPass := config.Admindb.Password
    dbName := config.Admindb.Dbname
	host := config.Admindb.Host
	port := strconv.Itoa(config.Admindb.Port)
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@("+host+":"+port+")/"+dbName)
	//db, err := sql.Open("mysql", "gitea:gitea@tcp(127.0.0.1:3306)/xqledgeradmindb")
	mainConn = db
	return err
}

func GetTenants() ([]utils.Tenant, error) {
	var result []utils.Tenant
	if mainConn == nil {
		err := dbConn()
		if err != nil {
			return result, err
		}
	}
	//defer mainConn.Close()
	var id int
	var subscription int64
	var active bool
	var name, description string
	rows, queryErr := mainConn.Query("SELECT TenantID,Name,Description,Subscription,Active FROM tenants")
	if queryErr != nil {
		return result, queryErr
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &subscription, &active)
		if err != nil {
			return result, err
		}
		dbs, dberr := getDBsByTenant(id)
		if dberr != nil {
			fmt.Println(dberr.Error())
			return result, dberr
		}
		t := utils.Tenant{TenantID: id, Name: name, Description: description, Subscription: subscription, Active: active, Databases: dbs}
		result = append(result, t)
	}
	return result, nil
}

func NewDB(name, description string, t utils.Tenant) error {
	if mainConn == nil {
		err := dbConn()
		if err != nil {
			return err
		}
	}
	now := getEpochTime()
	stmt, err := mainConn.Prepare("INSERT INTO databases(DatabaseID, Name, TenantID, Description, Creation, Active) VALUES (?,?,?,?,?,?)")
	if err != nil {
        return err
    }
    res, err := stmt.Exec(0, name, 0, description, now, true)
    if err != nil {
        return err
    }
    _, dberr := res.LastInsertId()
    if dberr != nil {
        return dberr
    }
	return nil
}

func getDBsByTenant(tenantID int) ([]utils.Database, error) {
	var result []utils.Database
	if mainConn == nil {
		err := dbConn()
		if err != nil {
			return result, err
		}
	}
	//defer mainConn.Close()
	var id int
	var creation int64
	var active bool
	var name, description string
	fmt.Println(tenantID)
	query := "SELECT DatabaseID, Name, Description, Creation, Active FROM xqledgeradmindb.`databases` WHERE TenantID = " + strconv.Itoa(tenantID)
	rows, queryErr := mainConn.Query(query)
	if queryErr != nil {
		return result, queryErr
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &creation, &active)
		if err != nil {
			return result, err
		}
		cols, colerr := getCollectionssByDB(id)
		if colerr != nil {
			fmt.Println(colerr.Error())
			return result, colerr
		}
		t := utils.Database{DatabaseID: id, Name: name, Description: description, Creation: creation, Active: active, Collections: cols}
		fmt.Println(t)
		result = append(result, t)
	}
	return result, nil
}

func getCollectionssByDB(databaseID int) ([]utils.Collection, error) {
	var result []utils.Collection
	if mainConn == nil {
		err := dbConn()
		if err != nil {
			return result, err
		}
	}
	//defer mainConn.Close()
	var id int
	var creation int64
	var active bool
	var name, description string
	//fmt.Println(databaseID)
	query := "SELECT CollectionID, Name, Description, Creation, Active FROM xqledgeradmindb.`collections` WHERE DatabaseID = " + strconv.Itoa(databaseID)
	rows, queryErr := mainConn.Query(query)
	if queryErr != nil {
		return result, queryErr
	}
	for rows.Next() {
		err := rows.Scan(&id, &name, &description, &creation, &active)
		if err != nil {
			return result, err
		}
		t := utils.Collection{CollectionID: id, Name: name, Description: description, Creation: creation, Active: active}
		fmt.Println(t)
		result = append(result, t)
	}
	return result, nil
}

func getCollectionsByID(cid int) (utils.Collection, error) {
	var result utils.Collection
	if mainConn == nil {
		err := dbConn()
		if err != nil {
			return result, err
		}
	}
	var id int
	var creation int64
	var name, description string
	var active bool
	query := "SELECT CollectionID,Name,Description,Creation,Active FROM xqledgeradmindb.`collections` WHERE CollectionID = " + strconv.Itoa(cid)
    err := mainConn.QueryRow(query).Scan(&id,&name,&description,&creation,&active)
	result.CollectionID = id
	result.Name = name
	result.Description = description
	result.Active = active
	result.Creation = creation
	if err != nil {
		return result, err
	}
	return result, nil
}

func NewSession(s utils.Session) (utils.Session, error) {
	var result utils.Session
	if mainConn == nil {
		err := dbConn()
		if err != nil {
			return result, err
		}
	}
	now := getEpochTime()
	stmt, err := mainConn.Prepare("INSERT INTO sessions(User, Description, StartTime, Branch, CollectionID) VALUES (?,?,?,?,?)")
	if err != nil {
        return result, err
    }
    res, err := stmt.Exec(s.User, s.Description, now, s.Branch, s.Collection.CollectionID)
    if err != nil {
        return result, err
    }
    lastId, err := res.LastInsertId()
    if err != nil {
        return result, err
    }
	col, colErr := getCollectionsByID(s.Collection.CollectionID)
	if colErr != nil {
        return result, colErr
    }
	result.SessionID = int64(lastId)
	result.Branch = s.Branch
	result.Description = s.Description
	result.EndTime = 0
	result.StartTime = now
	result.Collection = col
	result.User = s.User

	return result, nil
}

func CloseSession(id int64) error {
	now := getEpochTime()
	if mainConn == nil {
		err := dbConn()
		if err != nil {
			return  err
		}
	}
	stmt, stmtErr := mainConn.Prepare("UPDATE sessions set EndTime = ? where SessionID = ?")
    if stmtErr != nil {
		return stmtErr
	}
    _, execErr := stmt.Exec(now, id)
    if execErr != nil {
		return execErr
	}
	return nil
}

func GetAllSessionsByCollection(cid int) ([]utils.Session, error) {
	var result []utils.Session
	if mainConn == nil {
		err := dbConn()
		if err != nil {
			return result, err
		}
	}
	var id int64
	var collectionid int
	var starttime, endtime int64
	var user, description, branch string
	query := fmt.Sprintf("SELECT SessionID, User, Description, StartTime, EndTime, Branch, CollectionID FROM sessions WHERE CollectionID = %d", cid)
	rows, queryErr := mainConn.Query(query)
	if queryErr != nil {
		return result, queryErr
	}
	for rows.Next() {
		err := rows.Scan(&id, &user, &description, &starttime, &endtime, &branch, &collectionid)
		if err != nil {
			fmt.Println(1)
			return result, err
		}
		col, colErr := getCollectionsByID(collectionid)
		if colErr != nil {
			fmt.Println(2)
			return result, colErr
		}
		t := utils.Session{SessionID: id, User: user, Description: description, StartTime: starttime, EndTime: endtime, Branch: branch, Collection: col}
		fmt.Println(t)
		result = append(result, t)
	}
	return result, nil
}

func GetActiveSessionsByCollection(cid int) ([]utils.Session, error) {
	var result []utils.Session
	if mainConn == nil {
		err := dbConn()
		if err != nil {
			return result, err
		}
	}
	var id int64
	var collectionid int
	var starttime, endtime int64
	var user, description, branch string
	query := "SELECT SessionID, User, Description, StartTime, EndTime, Branch, CollectionID FROM xqledgeradmindb.`sessions` WHERE CollectionID = " + strconv.Itoa(cid) + " AND EndTime = 0"
	rows, queryErr := mainConn.Query(query)
	if queryErr != nil {
		return result, queryErr
	}
	for rows.Next() {
		err := rows.Scan(&id, &user, &description, &starttime, &endtime, &branch, &collectionid)
		if err != nil {
			return result, err
		}
		col, colErr := getCollectionsByID(collectionid)
		if colErr != nil {
			return result, colErr
		}
		t := utils.Session{SessionID: id, User: user, Description: description, StartTime: starttime, EndTime: endtime, Branch: branch, Collection: col}
		fmt.Println(t)
		result = append(result, t)
	}
	return result, nil
}

func GetSessionByID(sid int64) (utils.Session, error) {
	var result utils.Session
	if mainConn == nil {
		err := dbConn()
		if err != nil {
			return result, err
		}
	}
	var id int64
	var collectionid int
	var starttime, endtime int64
	var user, description, branch string
	query := "SELECT SessionID, User, Description, StartTime, EndTime, Branch, CollectionID FROM xqledgeradmindb.`sessions` WHERE SessionID = " + strconv.FormatInt(sid, 10)
    err := mainConn.QueryRow(query).Scan(&id, &user, &description, &starttime, &endtime, &branch, &collectionid)
	if err != nil {
		return result, err
	}
	col, colErr := getCollectionsByID(collectionid)
	if colErr != nil {
		return result, colErr
	}
	result = utils.Session{SessionID: id, User: user, Description: description, StartTime: starttime, EndTime: endtime, Branch: branch, Collection: col}
	return result, nil
}

func getEpochTime() int64 {
    return time.Now().Unix()
}

