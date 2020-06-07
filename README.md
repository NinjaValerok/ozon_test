***
Пара заданий от озона.
***
array_include - Сумма двух. 
Дана последовательность из целых положительных чисел. 
Необходимо записать в выходной файл "1", если в последовательности есть два числа сумма, которых равна значению "target" или "0" если таких нет.
Формат ввода
5

1 7 3 4 7 9

Формат вывода
1

***
chan - 2 канала.
Необходимо написать функцию func Merge2Channels(f func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int) в package main.
Описание ее работы:
n раз сделать следующее

прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
вычислить f(x1) + f(x2)
записать полученное значение в out
Функция Merge2Channels должна быть неблокирующей, сразу возвращая управление.
Функция f может работать долгое время, ожидая чего-либо или производя вычисления.