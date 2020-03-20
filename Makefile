
all: build run

build:
	go build -o covid

run:
	./covid

draw:
	curl http://localhost:8080/draw

test:
	curl http://localhost:8080/test


drawstate:
	curl http://localhost:8080/drawstate?state=AK
	sleep 5
	curl http://localhost:8080/drawstate?state=AL
	sleep 5
	curl http://localhost:8080/drawstate?state=AR
	sleep 5
	curl http://localhost:8080/drawstate?state=AZ
	sleep 5
	curl http://localhost:8080/drawstate?state=CA
	sleep 5
	curl http://localhost:8080/drawstate?state=CO
	sleep 5
	curl http://localhost:8080/drawstate?state=CT
	sleep 5
	curl http://localhost:8080/drawstate?state=DC
	sleep 5
	curl http://localhost:8080/drawstate?state=DE
	sleep 5
	curl http://localhost:8080/drawstate?state=FL
	sleep 5
	curl http://localhost:8080/drawstate?state=GA
	sleep 5
	curl http://localhost:8080/drawstate?state=HI
	sleep 5
	curl http://localhost:8080/drawstate?state=IA
	sleep 5
	curl http://localhost:8080/drawstate?state=ID
	sleep 5
	curl http://localhost:8080/drawstate?state=IL
	sleep 5
	curl http://localhost:8080/drawstate?state=IN
	sleep 5
	curl http://localhost:8080/drawstate?state=KS
	sleep 5
	curl http://localhost:8080/drawstate?state=KY
	sleep 5
	curl http://localhost:8080/drawstate?state=LA
	sleep 5
	curl http://localhost:8080/drawstate?state=MA
	sleep 5
	curl http://localhost:8080/drawstate?state=MD
	sleep 5
	curl http://localhost:8080/drawstate?state=ME
	sleep 5
	curl http://localhost:8080/drawstate?state=MI
	sleep 5
	curl http://localhost:8080/drawstate?state=MN
	sleep 5
	curl http://localhost:8080/drawstate?state=MO
	sleep 5
	curl http://localhost:8080/drawstate?state=MS
	sleep 5
	curl http://localhost:8080/drawstate?state=MT
	sleep 5
	curl http://localhost:8080/drawstate?state=NC
	sleep 5
	curl http://localhost:8080/drawstate?state=ND
	sleep 5
	curl http://localhost:8080/drawstate?state=NE
	sleep 5
	curl http://localhost:8080/drawstate?state=NH
	sleep 5
	curl http://localhost:8080/drawstate?state=NJ
	sleep 5
	curl http://localhost:8080/drawstate?state=NM
	sleep 5
	curl http://localhost:8080/drawstate?state=NV
	sleep 5
	curl http://localhost:8080/drawstate?state=NY
	sleep 5
	curl http://localhost:8080/drawstate?state=OH
	sleep 5
	curl http://localhost:8080/drawstate?state=OK
	sleep 5
	curl http://localhost:8080/drawstate?state=OR
	sleep 5
	curl http://localhost:8080/drawstate?state=PA
	sleep 5
	curl http://localhost:8080/drawstate?state=RI
	sleep 5
	curl http://localhost:8080/drawstate?state=SC
	sleep 5
	curl http://localhost:8080/drawstate?state=SD
	sleep 5
	curl http://localhost:8080/drawstate?state=TN
	sleep 5
	curl http://localhost:8080/drawstate?state=TX
	sleep 5
	curl http://localhost:8080/drawstate?state=UT
	sleep 5
	curl http://localhost:8080/drawstate?state=VA
	sleep 5
	curl http://localhost:8080/drawstate?state=VT
	sleep 5
	curl http://localhost:8080/drawstate?state=WA
	sleep 5
	curl http://localhost:8080/drawstate?state=WI
	sleep 5
	curl http://localhost:8080/drawstate?state=WV
	sleep 5
	curl http://localhost:8080/drawstate?state=WY
	sleep 5
	curl http://localhost:8080/drawstate?state=PR
	sleep 5
	curl http://localhost:8080/drawstate?state=VI
	sleep 5
	curl http://localhost:8080/drawstate?state=GU
	sleep 5
	curl http://localhost:8080/drawstate?state=MP
	sleep 5
	curl http://localhost:8080/drawstate?state=AS