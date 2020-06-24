create database suflex

create table if not exists public.usuario
(
	id serial not null
		constraint usuario_pk
			primary key,
	nome text,
	sobrenome text,
	username text,
	senha text,
	salt text,
	datacriacao timestamp
);

create unique index if not exists usuario_username_uindex
	on public.usuario (username);
