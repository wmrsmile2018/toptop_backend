/* constraint parent_table: users */
alter table places add constraint users_places foreign key (id_user) references users (id);

/* constraint parent_table: category */
alter table places add constraint category_places foreign key (id_category) references category (id);

/* constraint parent_table: filters */
alter table place_filters add constraint filters_places foreign key (id_filter) references filters (id);

/* constraint parent_table: user_tags */
alter table place_user_tags add constraint user_tags_place_user_tags foreign key (id_user_tags) references user_tags (id);

/* constraint parent_table: places */
alter table place_filters add constraint places_place_filters foreign key (id_place) references places (id);
alter table place_user_tags add constraint places_place_user_tags foreign key (id_place) references places (id);
alter table photos add constraint places_photos foreign key (id_place) references places (id);

