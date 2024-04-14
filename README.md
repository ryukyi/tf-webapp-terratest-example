# Overview
This repo explores possible ways of testing workspaces/workloads terraform infrastructure.

## Proof of Concept 1
* Create a terraform workspace of the infrastructure
* Broadly categorize tests so they include equal-to, contains or not empty
* leave open ended for more tailored go tests

## Proof of Concept 2  
* create similar to docs. 
* Timestamp resource group with unique test id so they don't destroy existing
* basic http status code test for webapp
