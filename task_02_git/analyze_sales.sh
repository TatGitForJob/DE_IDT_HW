#!/bin/bash

if [[ $# -ne 1 ]]; then
    printf "Need 1 argument: path to file with sales:\n./analyze_sales.sh ../sales.txt\n"
    exit 1
fi

FILE=$1

if [[ ! -f "$FILE" ]]; then
    echo "File $FILE not found"
    exit 1
fi

declare -A day_revenue
declare -A item_qty
declare -A item_sum
declare -A weekday_of_date

total_sum=0
count=0

while read -r date weekday item price qty extra; do
    count=$((count + 1))
    if [[ -n "$extra" ]]; then
        echo "more than 5 elements in line $count"
        exit 1
    fi

    if [[ -z "$date" || -z "$weekday" || -z "$item" || -z "$price" || -z "$qty" ]]; then
        echo "Line have less than 5 elements:"
        echo "$date $weekday $item $price $qty"
        exit 1
    fi

    revenue=$(echo "$price * $qty" | bc -l)
    total_sum=$(echo "$total_sum + $revenue" | bc -l)
    day_revenue["$date"]=$(echo "${day_revenue[$date]:-0} + $revenue" | bc -l)
    weekday_of_date["$date"]="$weekday"
    item_qty["$item"]=$(echo "${item_qty[$item]:-0} + $qty" | bc -l)
    item_sum["$item"]=$(echo "${item_sum[$item]:-0} + $revenue" | bc -l)
done < "$FILE"

best_day=""
best_day_sum=0

for d in "${!day_revenue[@]}"; do
    if [[ $(echo "${day_revenue[$d]} > $best_day_sum" | bc) -eq 1 ]]; then
        best_day="$d"
        best_day_sum="${day_revenue[$d]}"
    fi
done

best_weekday="${weekday_of_date[$best_day]}"


best_item=""
best_qty=0

for item in "${!item_qty[@]}"; do
    if (( item_qty[$item] > best_qty )); then
        best_item="$item"
        best_qty="${item_qty[$item]}"
    fi
done

best_item_sum="${item_sum[$best_item]}"

printf "Общая сумма продаж: %.2f\n" "$total_sum"
printf "День с наибольшей выручкой: %s %s (сумма продаж: %.2f)\n" "$best_day" "$best_weekday" "$best_day_sum"
printf "Популярный товар: %s (количество проданных единиц: %d, сумма продаж: %.2f)\n" "$best_item" "$best_qty" "$best_item_sum"