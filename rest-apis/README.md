Invoke-RestMethod -Uri http://localhost:8080/ -B "Colin" -Method POST

Invoke-WebRequest http://localhost:8080/ -B "Colin" -Method POST

//Pass params with Get
$Parameters = @{
  id = 'John'
  Children = 'Abe','Karen','Jo'
}
Invoke-RestMethod -Uri http://localhost:8080/ -B $Parameters -Method GET



// pass json 
$body = @"
 {
     "id": 100,
     "name": "tea",
     "description": "a nice cup of tea"
 }
"@

Invoke-WebRequest -Uri "http://localhost:8080/" -Method Post -Body $body -ContentType "application/json"


Invoke-WebRequest -Uri "http://localhost:8080/1" -Method Put 

$update = @"
 {
     "id": 1,
     "name": "tea",
     "description": "a nice cup of tea 2"
 }
"@

Invoke-WebRequest -Uri "http://localhost:8080/4" -Method Put -Body $update -ContentType "application/json"