PGDMP                         y         	   Ecommerce    13.3    13.3     ?           0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false            ?           0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false            ?           0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false            ?           1262    24576 	   Ecommerce    DATABASE     o   CREATE DATABASE "Ecommerce" WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE = 'English_United States.1252';
    DROP DATABASE "Ecommerce";
                postgres    false            ?            1259    24662 	   Customers    TABLE     ?   CREATE TABLE public."Customers" (
    customer_id integer,
    customer_name character varying(255),
    address character varying(255),
    phone character varying(255),
    email character varying(255),
    role integer
);
    DROP TABLE public."Customers";
       public         heap    postgres    false            ?            1259    24611    admin_account    TABLE     ?   CREATE TABLE public.admin_account (
    admin_id integer NOT NULL,
    name character varying(255) NOT NULL,
    username character varying(255) NOT NULL,
    password character varying(255) NOT NULL,
    role integer
);
 !   DROP TABLE public.admin_account;
       public         heap    postgres    false            ?            1259    24668    order_details    TABLE     ?   CREATE TABLE public.order_details (
    id integer,
    order_id integer,
    product_id integer,
    price integer,
    quantity integer
);
 !   DROP TABLE public.order_details;
       public         heap    postgres    false            ?            1259    24627    orders    TABLE     ?   CREATE TABLE public.orders (
    order_id integer NOT NULL,
    quantity integer NOT NULL,
    status integer NOT NULL,
    customer_id integer,
    address character varying(255)
);
    DROP TABLE public.orders;
       public         heap    postgres    false            ?            1259    24637    products    TABLE     B  CREATE TABLE public.products (
    product_id integer NOT NULL,
    product_name character varying(255) NOT NULL,
    price double precision NOT NULL,
    descriptions character varying(255) NOT NULL,
    image character varying(255) NOT NULL,
    categories character varying(255) NOT NULL,
    stock integer NOT NULL
);
    DROP TABLE public.products;
       public         heap    postgres    false            ?          0    24662 	   Customers 
   TABLE DATA           ^   COPY public."Customers" (customer_id, customer_name, address, phone, email, role) FROM stdin;
    public          postgres    false    203   ?       ?          0    24611    admin_account 
   TABLE DATA           Q   COPY public.admin_account (admin_id, name, username, password, role) FROM stdin;
    public          postgres    false    200          ?          0    24668    order_details 
   TABLE DATA           R   COPY public.order_details (id, order_id, product_id, price, quantity) FROM stdin;
    public          postgres    false    204   9       ?          0    24627    orders 
   TABLE DATA           R   COPY public.orders (order_id, quantity, status, customer_id, address) FROM stdin;
    public          postgres    false    201   V       ?          0    24637    products 
   TABLE DATA           k   COPY public.products (product_id, product_name, price, descriptions, image, categories, stock) FROM stdin;
    public          postgres    false    202          5           2606    24618     admin_account admin_account_pkey 
   CONSTRAINT     d   ALTER TABLE ONLY public.admin_account
    ADD CONSTRAINT admin_account_pkey PRIMARY KEY (admin_id);
 J   ALTER TABLE ONLY public.admin_account DROP CONSTRAINT admin_account_pkey;
       public            postgres    false    200            7           2606    24631    orders orders_pkey 
   CONSTRAINT     V   ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (order_id);
 <   ALTER TABLE ONLY public.orders DROP CONSTRAINT orders_pkey;
       public            postgres    false    201            9           2606    24644    products products_pkey 
   CONSTRAINT     \   ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (product_id);
 @   ALTER TABLE ONLY public.products DROP CONSTRAINT products_pkey;
       public            postgres    false    202            ?           6104    24585    customer    PUBLICATION     Q   CREATE PUBLICATION customer WITH (publish = 'insert, update, delete, truncate');
    DROP PUBLICATION customer;
                postgres    false            ?      x?????? ? ?      ?   (   x?3400?tL????L?zť??F?&?f??\1z\\\ ??y      ?      x?????? ? ?      ?      x?340?4?4???"?=...  z      ?   ?  x?ŕ]o?8???_???U??Q?ZM&?t???6?d??c?l?6??_?H?JiW??`0>??????DQ?5"?.?????JXĬ?T?`?|D?*??0U?V6??Ν Ӱb?d?8?W?v?W?ҭ?i[#m?t???N??ƽ%d?ۅF???!?-?!)??!?i???	M?8+t`?c??\k;???uVI?Y?t??`?[?18?!=????!?@{f%?A9m??????b??̇?%???L1?8b??;?ID???gE$?{2?t???p??8?P.
{?\?3?4???V:??c#E?0??wO˅???A???4?-??h,?|?r?#?T?m?? @????Gs??Ѽ?z;c?*]???x?geYy4???i?n?e:? ?g?7??RX!????B??%#1?)?	y????O????#?gKpv??:???$3=d#?,c?q?U??'7S?`H?97[J?8^$wQ??/ ?Qp?u͙?ыK¼????^X7?*?=??N??P#??9Mfx?JW???????y???????[??ma?`9΄?G???????ˢ?\?????I%????GKB???.?⋣??4Y??|?
?$??R(?ֺ(ē? ? q͚???T?h???9(+?7??/?1*?F?S\??z??{ Wb??J:$ݐ?or??e5tK?tV?}?Bs?a??Ϡ?m???>??_tBǭ4^+?.???BV??!?^]??Ǐ$???????F??L?rĽ??)?u9?'k???{???c?|?,b?g?Ч?????????M%=????0~^????x??????O?FO?d??ߐy?ȭ??Ln?????x?c?)??$،G?ׇwy???//?Fm???_[z?vz???X}?gw{?7??m????????x??cٜ??????vk?????K?k??t???[???uL'??d2?u??     