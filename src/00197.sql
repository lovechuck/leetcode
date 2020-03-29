/*197. 上升的温度*/
SELECT a.Id from Weather a,Weather b
where a.Temperature > b.Temperature
  and a.RecordDate = DATE_ADD(b.RecordDate,INTERVAL 1 DAY)
