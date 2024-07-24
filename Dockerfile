FROM golang:bullseye
WORKDIR /app
COPY * /app
#install yt-dlp
RUN apt update
RUN apt install python3 curl ffmpeg libsndfile1-dev alsa-utils -y
RUN curl -L https://github.com/yt-dlp/yt-dlp/releases/latest/download/yt-dlp -O
RUN chmod +x yt-dlp
RUN mv yt-dlp /bin/yt-dlp
RUN yt-dlp -U
#temp
RUN yt-dlp --quiet -f 'bestaudio' 'https://www.youtube.com/watch?v=69Km895ntNc' -o test
#install mpv
RUN curl --output-dir /etc/apt/trusted.gpg.d -O https://apt.fruit.je/fruit.gpg
RUN echo "deb http://apt.fruit.je/debian bullseye mpv" >> /etc/apt/sources.list.d/fruit.list
RUN apt update -y
RUN apt install mpv -y
RUN mpv --version
#ENTRYPOINT [ "tail","-f","/dev/null" ]
#Install mocyt
RUN make install
RUN mocyt --version
#docker build -t mocyt .
#docker run --rm -it -d --privileged=true --device=/dev/snd:/dev/snd --name mocyt mocyt
#docker exec -it mocyt bash



