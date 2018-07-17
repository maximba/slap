CREATE TABLE IF NOT EXISTS rooms (
  id INT NOT NULL AUTO_INCREMENT,
  name VARCHAR(80),
  PRIMARY KEY (id)
) ENGINE=INNODB;

CREATE TABLE IF NOT EXISTS turn (
  id INT NOT NULL AUTO_INCREMENT,
  room_id INT,
  priority TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  attendee VARCHAR(80),
  PRIMARY KEY (id),
  INDEX (room_id, priority, attendee),
  FOREIGN KEY (room_id) REFERENCES rooms(id) ON UPDATE CASCADE ON DELETE CASCADE
) ENGINE=INNODB;

USE slapdb

INSERT INTO rooms (name) VALUES ('Monkey Island');
INSERT INTO rooms (name) VALUES ('Gotham');
INSERT INTO rooms (name) VALUES ('New New York');
INSERT INTO rooms (name) VALUES ('Kitchen');
