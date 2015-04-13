FROM teemow/aqbanking

RUN sudo pacman -Sy && sudo pacman -S --noconfirm go 
RUN sudo pacman -S --noconfirm make gcc

# go-aqbanking expects aqbanking/banking.h under /usr/local
RUN sudo ln -s /usr/include/aqbanking5 /usr/local/include/aqbanking5
RUN sudo ln -s /usr/include/gwenhywfar4 /usr/local/include/gwenhywfar4

WORKDIR /app
ADD . /app

RUN sudo chown -R teemow.users /app
RUN make

ENTRYPOINT ["telebank"]
