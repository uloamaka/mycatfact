This is an API app for returning facts about my cat 

Task — Profile Endpoint (core requirements)
Required endpoint
Create a GET endpoint at: /me
The endpoint must return JSON data with Content-Type: application/json
Must integrate with the Cat Facts API to fetch dynamic cat facts
Response structure (required fields)
Your endpoint must return a JSON response in this exact format:
{
  "status": "success",
  "user": {
    "email": "<your email>",
    "name": "<your full name>",
    "stack": "<your backend stack>"
  },
  "timestamp": "<current UTC time in ISO 8601 format>",
  "fact": "<random cat fact from Cat Facts API>"
}