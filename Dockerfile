FROM golang:1.18

ARG WORK_DIR="/app"
ENV VOLUME_IN="${WORK_DIR}/input"
ENV VOLUME_OUT="${WORK_DIR}/output"

ENV input "paths.csv"
ENV output "result.csv"

WORKDIR ${WORK_DIR}

COPY . .

RUN go mod download
RUN go mod tidy

VOLUME ${WORK_DIR}/input
VOLUME ${WORK_DIR}/output

RUN make test
RUN make cli

CMD ["sh", "-c", "./bin/fare-cli -f ${VOLUME_IN}/${input} -o ${output}"]