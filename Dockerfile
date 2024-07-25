FROM golang:bullseye
WORKDIR /app
COPY . /app
#install yt-dlp
RUN apt update
RUN apt install python3 curl ffmpeg libsndfile1-dev alsa-utils -y
RUN curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -O
RUN chmod +x yt-dlp
RUN mv yt-dlp /bin/yt-dlp
RUN yt-dlp -U
#install mpv
RUN curl --output-dir /etc/apt/trusted.gpg.d -O https://apt.fruit.je/fruit.gpg
RUN echo "deb http://apt.fruit.je/debian bullseye mpv" >> /etc/apt/sources.list.d/fruit.list
RUN apt update -y
RUN apt install mpv -y
RUN mpv --version
#ENTRYPOINT [ "tail","-f","/dev/null" ]
#Install mocyt
RUN make install
#docker build -t mocyt .
#docker run --rm -it -d --privileged=true --device=/dev/snd:/dev/snd --name mocyt mocyt
#docker exec -it mocyt bash



