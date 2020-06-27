Invoke-RestMethod -Uri http://localhost:8080/ -B "Colin" -Method POST
Invoke-WebRequest http://localhost:8080/ -B "Colin" -Method POST

//Get
$Parameters = @{
  id = 'John'
  Children = 'Abe','Karen','Jo'
}
Invoke-RestMethod -Uri http://localhost:8080/ -B $Parameters -Method GET


//POST (ADD) 
$body = @"
 {
     "id": 100,
     "name": "tea",
     "description": "a nice cup of tea"
 }
"@
// write data with POST request
Invoke-WebRequest -Uri "http://localhost:8080/" -Method Post -Body $body -ContentType "application/json"

//PUT (UPDATE)
$update = @"
 {
     "id": 1,
     "name": "tea",
     "description": "a nice cup of tea 2"
 }
"@
Invoke-WebRequest -Uri "http://localhost:8080/4" -Method Put -Body $update -ContentType "application/json"

//POST (Test Validation)
//POST (ADD) 
$bodyval = @"
 {
     "id": 100,
     "name": "tea",
     "description": "a nice cup of tea",
     "price": 1.0,
     "sku": "abc-asd-def"
 }
"@
// write data with POST request
Invoke-WebRequest -Uri "http://localhost:8080/" -Method Post -Body $bodyval -ContentType "application/json"
