drop database if exists EventosDB;

create database EventosDB;
 
\c eventosdb;

drop TABLE IF EXISTS Eventos;
CREATE TABLE IF NOT EXISTS Eventos(
ID_Evento SERIAL,
Nombre varchar(30),
Descripcion varchar(30),
HoraDeInicio varchar(30),
HoraDeFinalizacion varchar(30),
CONSTRAINT PK_Evento PRIMARY KEY (ID_Evento));
