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
     "id": 1,
     "name": "tea",
     "description": "a nice cup of tea"
 }
"@

Invoke-WebRequest -Uri "http://localhost:8080/" -Method Post -Body $body -ContentType "application/json"