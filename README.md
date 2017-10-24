# gett

So, I picked Golang. Most of the time I use Ruby in my work. I’m just starting with Go and don’t fully understand conventions on how to write and organize code yet and I guess that is why I had some troubles splitting the code onto multiple packages and testing it.

Some notes:

- I picked some packages like gorm and echo to build this api. After I’ve implemented working solution with them and started to work on splitting the project I thought I would have been probably okay with net/http and database/sql at least for such api. Ruby way is to use gems before you start implementing your own solution for a problem. Not sure about Golang. Maybe it’s preferable to use standard libs.

- routes. I would have been preferred “/drivers/1” and “/drivers/import”. They seem more restful. I decided to keep the ones described in the challenge because, I guess, you expect them to be like that.

- “/import” route behavior depends heavily on business logic. There are so many ways this route can behave, for example: a) reject request if at least 1 record is invalid/exists b) partially save, e.g. 7/10 drivers are saved (maybe use some rare status codes like 207 multi-status. If record is already exists should you update it or not?(something like upsert) etc. I decided to just accept payload if it’s valid and respond with 202 status accepted, because other implementation need reasoning to justify them. If record exists - I skip it. If not - save it. If invalid - print error.

- talking about validation: I decided to set uniqueness validation on license number column because that seemed reasonable. Not every driver from your input JSON will pass this rule. :)

- deployed to DO with docker.
