CREATE TABLE IF NOT EXISTS agenda.contacto
(
    id serial NOT NULL,
    nombre character varying(255) NOT NULL,
    documento character varying(10) NOT NULL,
    direccion character varying(255) NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL,
    fecha_modificacion timestamp without time zone NOT NULL,
    CONSTRAINT pk_contacto PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS agenda.telefono
(
    id serial NOT NULL,
    telefono character varying(10) NOT NULL,
    contacto_id integer NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL,
    fecha_modificacion timestamp without time zone NOT NULL,
    CONSTRAINT pk_telefono PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS agenda.correo
(
    id serial NOT NULL,
    correo character varying(255) NOT NULL,
    contacto_id integer NOT NULL,
    activo boolean NOT NULL DEFAULT true,
    fecha_creacion timestamp without time zone NOT NULL,
    fecha_modificacion timestamp without time zone NOT NULL,
    CONSTRAINT pk_correo PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS agenda.telefono
    ADD CONSTRAINT fk_telefono_contacto FOREIGN KEY (contacto_id)
    REFERENCES agenda.contacto (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE CASCADE
    NOT VALID;


ALTER TABLE IF EXISTS agenda.correo
    ADD CONSTRAINT fk_correo_contacto FOREIGN KEY (contacto_id)
    REFERENCES agenda.contacto (id) MATCH SIMPLE
    ON UPDATE CASCADE
    ON DELETE CASCADE
    NOT VALID;

END;
