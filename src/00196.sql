
delete a from Person a, (select min(Id), Email from Person group by Email having count(Email) > 1 order by id desc) b
where a.Id <> b.Id and a.Email = b.Email

delete a from Person a,Person b
where a.Email = b.Email and a.Id > b.Id