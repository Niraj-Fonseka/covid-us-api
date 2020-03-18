
all: build run

build:
	go build -o covid

run:
	./covid

draw:
	curl http://localhost:8080/draw


drawstate:
	curl http://localhost:8080/drawstate?state=AK
	curl http://localhost:8080/drawstate?state=AL
	curl http://localhost:8080/drawstate?state=AR
	curl http://localhost:8080/drawstate?state=AZ
	curl http://localhost:8080/drawstate?state=CA
	curl http://localhost:8080/drawstate?state=CO
	curl http://localhost:8080/drawstate?state=CT
	curl http://localhost:8080/drawstate?state=DC
	curl http://localhost:8080/drawstate?state=DE
	curl http://localhost:8080/drawstate?state=FL
	curl http://localhost:8080/drawstate?state=GA
	curl http://localhost:8080/drawstate?state=HI
	curl http://localhost:8080/drawstate?state=IA
	curl http://localhost:8080/drawstate?state=ID
	curl http://localhost:8080/drawstate?state=IL
	curl http://localhost:8080/drawstate?state=IN
	curl http://localhost:8080/drawstate?state=KS
	curl http://localhost:8080/drawstate?state=KY
	curl http://localhost:8080/drawstate?state=LA
	curl http://localhost:8080/drawstate?state=MA
	curl http://localhost:8080/drawstate?state=MD
	curl http://localhost:8080/drawstate?state=ME
	curl http://localhost:8080/drawstate?state=MI
	curl http://localhost:8080/drawstate?state=MN
	curl http://localhost:8080/drawstate?state=MO
	curl http://localhost:8080/drawstate?state=MS
	curl http://localhost:8080/drawstate?state=MT
	curl http://localhost:8080/drawstate?state=NC
	curl http://localhost:8080/drawstate?state=ND
	curl http://localhost:8080/drawstate?state=NE
	curl http://localhost:8080/drawstate?state=NH
	curl http://localhost:8080/drawstate?state=NJ
	curl http://localhost:8080/drawstate?state=NM
	curl http://localhost:8080/drawstate?state=NV
	curl http://localhost:8080/drawstate?state=NY
	curl http://localhost:8080/drawstate?state=OH
	curl http://localhost:8080/drawstate?state=OK
	curl http://localhost:8080/drawstate?state=OR
	curl http://localhost:8080/drawstate?state=PA
	curl http://localhost:8080/drawstate?state=RI
	curl http://localhost:8080/drawstate?state=SC
	curl http://localhost:8080/drawstate?state=SD
	curl http://localhost:8080/drawstate?state=TN
	curl http://localhost:8080/drawstate?state=TX
	curl http://localhost:8080/drawstate?state=UT
	curl http://localhost:8080/drawstate?state=VA
	curl http://localhost:8080/drawstate?state=VT
	curl http://localhost:8080/drawstate?state=WA
	curl http://localhost:8080/drawstate?state=WI
	curl http://localhost:8080/drawstate?state=WV
	curl http://localhost:8080/drawstate?state=WY
	curl http://localhost:8080/drawstate?state=PR
	curl http://localhost:8080/drawstate?state=VI
	curl http://localhost:8080/drawstate?state=GU
	curl http://localhost:8080/drawstate?state=MP
	curl http://localhost:8080/drawstate?state=AS