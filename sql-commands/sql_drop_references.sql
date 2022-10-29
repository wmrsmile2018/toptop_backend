alter table places drop constraint users_places;
alter table places drop constraint category_places;
alter table place_filters drop constraint filters_places;
alter table place_user_tags drop constraint user_tags_place_user_tags;
alter table place_filters drop constraint places_place_filters;
alter table place_user_tags drop constraint places_place_user_tags;
alter table photos drop constraint places_photos;