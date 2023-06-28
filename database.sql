CREATE DATABASE project_go;

use project_go;

CREATE TABLE users(
                      id integer auto_increment,
                      name varchar(70) not null,
                      email varchar(100) not null,
                      no_hp varchar(100) not null,
                      password varchar(100) not null
                      primary key(id)
)engine = innoDB;

CREATE TABLE cars(
                     id integer auto_increment,
                     title varchar(100) not null,
                     price integer not null,
                     image varchar(255) not null ,
                     description text not null,
                     passenger tinyint not null ,
                     luggage tinyint not null ,
                     car_type varchar(100) not null,
                     isDriver boolean not null,
                     primary key(id)
)engine = innoDB;

CREATE TABLE car_rating(
                           id integer auto_increment,
                           user_id integer not null,
                           car_id integer not null,
                           primary key(id),
                           CONSTRAINT fk_user_rating
                               FOREIGN KEY (user_id) REFERENCES users(id)
                                   ON DELETE CASCADE ,
                           CONSTRAINT fk_car_rating
                               FOREIGN KEY (car_id) REFERENCES cars(id)
                                   ON DELETE CASCADE
) ENGINE = innoDB;

CREATE TABLE tours(
                      id  integer not null auto_increment,
                      title varchar(100) not null ,
                      price integer not null ,
                      duration varchar(100) not null,
                      description text not null ,
                      primary key (id)
) engine = innoDB;

CREATE TABLE car_tours(
                          id  integer not null auto_increment,
                          car_id integer not null ,
                          tour_id integer not null ,
                          primary key (id),
                          CONSTRAINT fk_car_id FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE,
                          CONSTRAINT fk_tour_id FOREIGN KEY (tour_id) REFERENCES tours(id) ON DELETE CASCADE
) engine = innoDB;

CREATE TABLE lease_types(
                            id integer auto_increment not null,
                            title varchar(100) not null ,
                            description text not null ,
                            primary key (id)
) engine = innoDB;

CREATE TABLE car_lease_types(
                                id integer auto_increment not null,
                                car_id integer not null ,
                                lease_type_id integer not null,
                                primary key (id),
                                CONSTRAINT fk_car_lease FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE,
                                CONSTRAINT fk_lease_type_id FOREIGN KEY (lease_type_id) REFERENCES lease_types(id) ON DELETE CASCADE
) engine = innoDB;

CREATE TABLE transaction(
                            id integer auto_increment not null,
                            car_id integer not null ,
                            user_id integer not null,
                            price integer not null,
                            primary key (id),
                            CONSTRAINT fk_car_transaction FOREIGN KEY (car_id) REFERENCES cars(id) ON DELETE CASCADE,
                            CONSTRAINT fk_user_transaction FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) engine = innoDB;