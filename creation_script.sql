CREATE TABLE rooms(
                           ID  serial PRIMARY KEY ,
                           name CHAR(255) NOT NULL,
);
CREATE TABLE sensors(
                       ID  serial PRIMARY KEY ,
                       name CHAR(255) NOT NULL,
                       roomID INT NOT NULL
);

CREATE TABLE presences(
                       ID  serial PRIMARY KEY ,
                       MAC CHAR(255) NOT NULL,
                       lastDetected TIMESTAMP NOT NULL,
                       active BOOLEAN NOT NULL,
                       roomID INT NOT NULL
);

ALTER TABLE sensors
    ADD CONSTRAINT roomID
        FOREIGN KEY (roomID)
            REFERENCES rooms (ID);

ALTER TABLE presences
    ADD CONSTRAINT roomID
        FOREIGN KEY (roomID)
            REFERENCES rooms (ID);