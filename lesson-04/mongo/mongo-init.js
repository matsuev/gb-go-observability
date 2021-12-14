db.auth('alex', 'pass')

db.users.drop(),
db.users.insert({ uid: 1, lname: 'Matsuev', fname: 'Alex', email: 'alex.matsuev@gmail.com' }),
db.users.insert({ uid: 2, lname: 'Ivanov', fname: 'Ivan', email: 'ivanov@example.com' })
