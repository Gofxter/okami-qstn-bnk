create table IF NOT EXISTS questions(
                          Id UUID PRIMARY KEY,
                          Role varchar(25) NOT NULL,
                          Topic varchar(100) NOT NULL,
                          Type varchar(20) NOT NULL,
                          Difficulty smallint NOT NULL,
                          Text text NOT NULL
);

create table IF NOT EXISTS options(
                                      Id UUID PRIMARY KEY,
                                      Question_Id UUID NOT NULL,
                                      Text text NOT NULL,
                                      Is_Correct bool NOT NULL,
                                      FOREIGN KEY (Question_Id) REFERENCES questions(Id)
);

create table IF NOT EXISTS test_templates(
                          Id UUID PRIMARY KEY,
                          Role varchar(25) NOT NULL,
                          Purpose varchar(50) NOT NULL

);
