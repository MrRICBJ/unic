Create table users (
                       id BIGSERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       password VARCHAR(255) NOT NULL,
                       role VARCHAR(255) NOT NULL,
                       result_test VARCHAR(255)[],
                       students INT[],
);

CREATE table data (
    theory Text,
    test Text[]
);

INSERT INTO data (theory, test)
VALUES ('SQL (Structured Query Language) - это язык программирования, который используется для
управления базами данных. Первый прототип языка SQL представила в 1979 году компания-
разработчик Oracle. Сначала это был простейший инструмент для извлечения нужных данных,
вроде фильтров в Excel-таблицах. С годами он усложнился, и теперь его применяют в качестве
одного из основных инструментов для обработки данных. С помощью SQL можно:
● собирать и хранить данные в виде таблиц;
● изменять их содержимое и структуру;
● объединять данные и выполнять вычисления;
● защищать и распределять доступ.
Например, в компании работает 500 сотрудников. 100 из них занимаются продажами и постоянно
пользуются CRM, чтобы вносить данные о клиентах: новые договоры, суммы, скидки и контакты. И
есть 15 IT-специалистов, которые настраивают, обновляют и меняют структуру CRM, когда это
требуется. 20 сотрудников бухгалтерии регулярно выгружают из системы данные об оплатах,
выставленные счета и акты на подпись. При помощи SQL можно предоставить всем им доступ
только к нужной части CRM, чтобы никто случайно не повредил важные данные или элементы
кода.
Особенности языка SQL
● Это язык запросов, а не программирования. Его используют в дополнение к Python, JavaScript
или C++, но только для работы с базами данных. Написать на нём полноценный сайт или
приложение невозможно.
● Чёткая и понятная структура. Это делает язык SQL для работы с данными относительно простым
для начинающих.
● Универсальность. Есть единые стандарты построения запросов для любых баз данных и
браузеров, которые позволяют обрабатывать даже очень большие объёмы информации.
● Совместный доступ. SQL позволяет создавать интерактивные запросы. Это значит, что можно
получать нужные данные онлайн и принимать решения на их основе.
● Управление доступом. При помощи SQL можно предоставить, закрыть или ограничить доступ к
данным для разных групп пользователей, а также выдать им определённый набор функций:
чтение, изменение, создание, удаление, копирование. Это защищает базы данных от взлома или
несогласованных изменений.
Структура базы данных SQL состоит из шести элементов:
● Ключ — самый главный столбец, который связывает таблицы данных.
      Они бывают:
      - первичными — содержат уникальный идентификатор для каждого объекта, например,
артикул;
      - потенциальными — содержат альтернативный идентификатор;
      - внешними — содержат ссылку, которая позволяет связать две таблицы, при этом значения
ключей в одной таблице соответствуют первичному ключу в другой.

● Отношение — таблица с данными, представленными в строках и столбцах.
● Атрибут — столбец, который содержит наименование, тип, цену или другой параметр.
● Домен — значения, которые допустимы для данного атрибута: например, стоимость в рублях
или название кириллическими символами.
● Кортеж — пронумерованная строка, где содержатся все данные о каком-либо объекте.
● Значение — содержимое ячейки в таблице.
У столбца есть название и тип данных. Команда создания таблицы должна содержать все
вышеупомянутое:
CREATE TABLE table_name
(
    column_name_1 column_type_1,
    column_name_2 column_type_2,
    ...,
    column_name_N column_type_N,
);

table_name — имя таблицы;
column_name — имя столбца;
column_type — тип данных столбца.
Имена могут содержать символы подчеркивания для большей наглядности. Классический пример
непонятных названий — table1, table2 и т. п. Использование транслита, неясных сокращений и,
разумеется, наличие орфографических ошибок тоже не приветствуется. Хороший пример коротких
информативных названий: Customers, Users, Orders, так как по названию таблицы должно быть
очевидно, какие данные таблица будет содержать. Эта же логика применима и к названию
столбцов. Максимальная длина названия и для таблицы, и для столбцов — 64 символа.
Для каждого столбца таблицы будет определен тип данных. Неправильное использование типов
данных увеличивает как объем занимаемой памяти, так и время выполнения запросов к таблице.
Это может быть незаметно на таблицах в несколько строк, но очень существенно, если количество
строк будет измеряться десятками и сотнями тысяч, и это далеко не предел для рабочей базы
данных.
Числовые типы
INT — целочисленные значения от −2147483648 до 2147483647, 4 байта.
DECIMAL — хранит числа с заданной точностью. Использует два параметра — максимальное
количество цифр всего числа (precision) и количество цифр дробной части (scale). Рекомендуемый
тип данных для работы с валютами и координатами. Можно использовать синонимы NUMERIC,
DEC, FIXED.
TINYINT — целые числа от −127 до 128, занимает 1 байт хранимой памяти.
BOOL — 0 или 1. Однозначный ответ на однозначный вопрос — false или true. Название столбцов
типа boolean часто начинается с is, has, can, allow. По факту это даже не отдельный тип данных, а
псевдоним для типа TINYINT (1). Тип настолько востребован на практике, что для него в MySQL
создали встроенные константы FALSE (0) или TRUE (1). Можно использовать синоним BOOLEAN.
FLOAT — дробные числа с плавающей запятой (точкой).
Символьные
VARCHAR(N) — N определяет максимально возможную длину строки. Создан для хранения
текстовых данных переменной длины, поэтому память хранения зависит от длины строки.
Наиболее часто используемый тип строковых данных.
CHAR(N) — как и с varchar, N указывает максимальную длину строки. Char создан хранить данные
строго фиксированной длины, и каждая запись будет занимать ровно столько памяти, сколько
требуется для хранения строки длиной N.
TEXT — подходит для хранения большого объема текста до 65 KB, например, целой статьи.
Дата и время
DATE — только дата. Диапазон от 1000-01-01 по 9999-12-31. Подходит для хранения дат рождения,
исторических дат, начиная с 11 века. Память хранения — 3 байта.
TIME — только время — часы, минуты, секунды — «hh:mm:ss». Память хранения — 3 байта.
DATETIME — соединяет оба предыдущих типа — дату и время. Использует 8 байтов памяти.
TIMESTAMP — хранит дату и время начиная с 1970 года. Подходит для большинства бизнес-задач.
Потребляет 4 байта памяти, что в два раза меньше, чем DATETIME, поскольку использует более
скромный диапазон дат.
Бинарные
Используются для хранения файлов, фото, документов, аудио и видеоконтента. Все это хранится в
бинарном виде.
BLOB — до 65 КБ бинарных данных
LARGEBLOB — до 4 ГБ.
INSERT
Самый простой способ вставить данные в таблицу — команда INSERT INTO table_name
VALUES(column1, column2, column_n). Допустим, что дана таблица users с тремя полями name, age
и contact, а также типами varchar, integer и varchar. Применим следующую команду для вставки
данных в таблицу users:
INSERT INTO users VALUES("Brave Zombie",18,"987654321");
SELECT
Общая структура запроса выглядит следующим образом:
SELECT (''столбцы или * для выбора всех столбцов; обязательно'')
FROM (''таблица; обязательно'')
WHERE (''условие/фильтрация, например, city = ''Moscow''; необязательно'')
GROUP BY (''столбец, по которому хотим сгруппировать данные; необязательно'')
HAVING (''условие/фильтрация на уровне сгруппированных данных; необязательно'')
ORDER BY (''столбец, по которому хотим отсортировать вывод; необязательно'')
Разберем структуру. Для удобства текущий изучаемый элемент в запроса выделяется CAPS''ом.
SELECT, FROM
SELECT, FROM — обязательные элементы запроса, которые определяют выбранные столбцы, их
порядок и источник данных.
Выбрать все (обозначается как *) из таблицы Customers:
SELECT * FROM Customers
Выбрать столбцы CustomerID, CustomerName из таблицы Customers:
SELECT CustomerID, CustomerName FROM Customers
WHERE
WHERE — необязательный элемент запроса, который используется, когда нужно отфильтровать
данные по нужному условию. Очень часто внутри элемента where используются IN / NOT IN для
фильтрации столбца по нескольким значениям, AND / OR для фильтрации таблицы по нескольким
столбцам.
Фильтрация по одному условию и одному значению:
select * from Customers
WHERE City = ''London''

Фильтрация по одному условию и нескольким значениям с применением IN (включение) или NOT
IN (исключение):
select * from Customers
where City IN (''London'', ''Berlin'')
select * from Customers
where City NOT IN (''Madrid'', ''Berlin'',''Bern'')
Фильтрация по нескольким условиям с применением AND (выполняются все условия) или OR
(выполняется хотя бы одно условие) и нескольким значениям:
select * from Customers
where Country = ''Germany'' AND City not in (''Berlin'', ''Aachen'') AND CustomerID > 15
select * from Customers
where City in (''London'', ''Berlin'') OR CustomerID > 4
GROUP BY
GROUP BY — необязательный элемент запроса, с помощью которого можно задать агрегацию по
нужному столбцу (например, если нужно узнать какое количество клиентов живет в каждом из
городов).
При использовании GROUP BY обязательно:
перечень столбцов, по которым делается разрез, был одинаковым внутри SELECT и внутри GROUP
BY,
агрегатные функции (SUM, AVG, COUNT, MAX, MIN) должны быть также указаны внутри SELECT с
указанием столбца, к которому такая функция применяется.
Группировка количества клиентов по городу:
select City, count(CustomerID) from Customers
GROUP BY City
Группировка количества клиентов по стране и городу:
select Country, City, count(CustomerID) from Customers
GROUP BY Country, City
Группировка продаж по ID товара с разными агрегатными функциями: количество заказов с
данным товаром и количество проданных штук товара:
select ProductID, COUNT(OrderID), SUM(Quantity) from OrderDetails
GROUP BY ProductID
Группировка продаж с фильтрацией исходной таблицы. В данном случае на выходе будет таблица
с количеством клиентов по городам Германии:
select City, count(CustomerID) from Customers
WHERE Country = ''Germany''
GROUP BY City
Переименование столбца с агрегацией с помощью оператора AS. По умолчанию название столбца
с агрегацией равно примененной агрегатной функции, что далее может быть не очень удобно для
восприятия.
select City, count(CustomerID) AS Number_of_clients from Customers
group by City
HAVING
HAVING — необязательный элемент запроса, который отвечает за фильтрацию на уровне
сгруппированных данных (по сути, WHERE, но только на уровень выше).
Фильтрация агрегированной таблицы с количеством клиентов по городам, в данном случае
оставляем в выгрузке только те города, в которых не менее 5 клиентов:
select City, count(CustomerID) from Customers
group by City
HAVING count(CustomerID) >= 5
В случае с переименованным столбцом внутри HAVING можно указать как и саму агрегирующую
конструкцию count(CustomerID), так и новое название столбца number_of_clients:
select City, count(CustomerID) as number_of_clients from Customers
group by City
HAVING number_of_clients >= 5
Пример запроса, содержащего WHERE и HAVING. В данном запросе сначала фильтруется исходная
таблица по пользователям, рассчитывается количество клиентов по городам и остаются только те
города, где количество клиентов не менее 5:
select City, count(CustomerID) as number_of_clients from Customers
WHERE CustomerName not in (''Around the Horn'',''Drachenblut Delikatessend'')
group by City
HAVING number_of_clients >= 5
ORDER BY
ORDER BY — необязательный элемент запроса, который отвечает за сортировку таблицы.
Простой пример сортировки по одному столбцу. В данном запросе осуществляется сортировка по
городу, который указал клиент:
select * from Customers
ORDER BY City
Осуществлять сортировку можно и по нескольким столбцам, в этом случае сортировка происходит
по порядку указанных столбцов:
select * from Customers
ORDER BY Country, City
По умолчанию сортировка происходит по возрастанию для чисел и в алфавитном порядке для
текстовых значений. Если нужна обратная сортировка, то в конструкции ORDER BY после названия
столбца надо добавить DESC:
select * from Customers
order by CustomerID DESC
Обратная сортировка по одному столбцу и сортировка по умолчанию по второму:
select * from Customers
order by Country DESC, City
JOIN
JOIN — необязательный элемент, используется для объединения таблиц по ключу, который
присутствует в обеих таблицах. Перед ключом ставится оператор ON.
Запрос, в котором соединяем таблицы Order и Customer по ключу CustomerID, при этом перед
названиям столбца ключа добавляется название таблицы через точку:
select * from Orders
JOIN Customers ON Orders.CustomerID = Customers.CustomerID
Нередко может возникать ситуация, когда надо промэппить одну таблицу значениями из другой.
В зависимости от задачи, могут использоваться разные типы присоединений. INNER JOIN —
пересечение, RIGHT/LEFT JOIN для мэппинга одной таблицы знаениями из другой,
select * from Orders
join Customers on Orders.CustomerID = Customers.CustomerID
where Customers.CustomerID >10
UPDATE
UPDATE — это команда, которая обновляет данные в таблице. Ее общий синтаксис такой:
UPDATE [table] table_name
SET column1 = value1, column2 = value2, ...
[WHERE condition]
[ORDER BY expression [ ASC | DESC ]]
[LIMIT number_rows];
Сначала мы указываем обязательные параметры: название таблицы, названия колонок и нужные
значения для обновления. Обратите внимание, что в MySQL можно использовать ключевое слово
table (update table), а можно его опустить и сразу указать название таблицы.
Затем идут необязательные блоки WHERE (условие обновления), ORDER BY (сортировка) и LIMIT
(ограничение количества обновляемых записей).
DELETE
Оператор DELETE удаляет строки из временных или постоянных базовых таблиц, представлений
или курсоров, причем в двух последних случаях действие оператора распространяется на те
базовые таблицы, из которых извлекались данные в эти представления или курсоры. Оператор
удаления имеет простой синтаксис:
DELETE FROM <имя таблицы >
[WHERE <предикат>];
Если предложение WHERE отсутствует, удаляются все строки из таблицы или представления
(представление должно быть обновляемым).',
        ARRAY['Какие задачи можно выполнить в SQL?',
        'Какие существуют типы данных в SQL?',
        'Какие из перечисленных типов данных не относятся к вещественным типам:',
        'Что такое бинарные типы данных?',
        'Какой оператор SQL используется для вставки новых данных в базу данных?',
        'Для создания новой таблицы в существующей базе данных используют команду:',
        'Имеются элементы запроса: 1. SELECT employees.name, departments.name; 2. ON
employees.department_id=departments.id; 3. FROM employees; 4. LEFT JOIN departments. В каком
порядке их нужно расположить, чтобы выполнить поиск имен всех работников со всех отделов?',
        'Как расшифровывается SQL?',
        'Запрос для выборки всех значений из таблицы «Persons» имеет вид:',
        'Напишите запрос, возвращающий значения из колонки «FirstName» таблицы «Users».',
        'Что возвращает запрос SELECT * FROM Students?',
        'Обязательными фразами в запросе на выборку данных являются:',
        'Что будет если не использовать элемент WHERE в команде DELETE?']);