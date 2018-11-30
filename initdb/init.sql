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

INSERT INTO rooms (name) VALUES ('Monza');
INSERT INTO rooms (name) VALUES ('Marina Bay');
INSERT INTO rooms (name) VALUES ('Montmelo');
INSERT INTO rooms (name) VALUES ('Silverstone');
INSERT INTO rooms (name) VALUES ('Suzuka');
INSERT INTO rooms (name) VALUES ('Hockenheim');
INSERT INTO rooms (name) VALUES ('Indianapolis');
INSERT INTO rooms (name) VALUES ('Kitchen');

