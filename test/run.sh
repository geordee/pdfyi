#!/usr/bin/env bash

printf "\n#1: Check Service\n"
curl -s -H "Authorization: Bearer Token" http://localhost:9090

printf "\n#2: Generate PDF\n"
curl -s -H "Authorization: Bearer Token" http://localhost:9090/pdfs -d @test/a4.html -o test/a4.pdf
