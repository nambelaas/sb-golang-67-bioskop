CREATE TABLE
    bioskop (
        Id serial primary key,
        Nama varchar(100) not null,
        Lokasi varchar(255) not null,
        Rating numeric (10, 2)
    );