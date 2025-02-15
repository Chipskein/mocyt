FROM golang:bullseye
WORKDIR /app
COPY . /app

RUN apt update
RUN apt install python3 curl ffmpeg libsndfile1-dev alsa-utils pulseaudio -y
RUN curl --output-dir /etc/apt/trusted.gpg.d -O https://apt.fruit.je/fruit.gpg
RUN echo "deb http://apt.fruit.je/debian bullseye mpv" >> /etc/apt/sources.list.d/fruit.list
RUN apt update -y
RUN apt install mpv -y
RUN mpv --version
RUN curl -L "https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp" -O
RUN chmod +x yt-dlp
RUN mv yt-dlp /bin/yt-dlp
RUN yt-dlp -U
RUN go build -o mocyt main.go
RUN cp mocyt /usr/bin/mocyt
ENTRYPOINT  ["tail","-f","/dev/null"]

#docker build -t chipskein/mocyt .
#docker run --rm -it -d --privileged=true -v /dev/snd:/dev/snd -e PULSE_SERVER=unix:${XDG_RUNTIME_DIR}/pulse/native -v ${XDG_RUNTIME_DIR}/pulse/native:${XDG_RUNTIME_DIR}/pulse/native -v ~/.config/pulse/cookie:/root/.config/pulse/cookie --name mocyt chipskein/mocyt
#docker exec -it mocyt bash


