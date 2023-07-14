INSERT INTO user (email, password_hash, name, role)
VALUES ('%s',' %s', '%s', '%s');
SELECT currval(pg_get_serial_sequence('user','id'));