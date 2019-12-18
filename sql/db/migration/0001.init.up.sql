CREATE TABLE Contacts (
    ID int(10) AUTO_INCREMENT,
    first_name varchar(30) NOT NULL,
    last_name  varchar(30) NOT NULL,
    phone      varchar(30) NOT NULL UNIQUE Constrain,
    email      varchar(30) NOT NULL UNIQUE Constrain,
);
