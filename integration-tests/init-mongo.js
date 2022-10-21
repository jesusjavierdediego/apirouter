db.createUser(
    {
      user: "xqledger",
      pwd: "toor",
      roles: [ "root"]
      // roles: [ "readWrite"]
    }
 );
 db.createCollection("main");
//  db.subscriptions.insertOne({ id: "1", email: "test@test.com", date_of_birth: "12-12-1970", consent: true, newsletter_id: "15", first_name: "John", gender: "Male", error: ""});
 