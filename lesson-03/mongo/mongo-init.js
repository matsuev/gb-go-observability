db.auth('alex', 'pass')

db.users.drop(),
db.users.insert({ lname: 'Matsuev', fname: 'Alex', email: 'alex.matsuev@gmail.com' }),
db.users.insert({ lname: 'Ivanov', fname: 'Ivan', email: 'ivanov@example.com' })
