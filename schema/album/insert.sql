INSERT INTO album (artist, price, title)
VALUES ('%s',' %d', '%s');
SELECT currval(pg_get_serial_sequence('album','id'));