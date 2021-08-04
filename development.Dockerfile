FROM golang:1.17rc2

ARG USERNAME=development
ARG WORKDIR=/home/$USERNAME

RUN useradd -ms /bin/bash $USERNAME

COPY [".","$WORKDIR"]

RUN find "$WORKDIR/" -type d -exec chmod 755 {} \;
RUN find "$WORKDIR/" -type f -exec chmod 755 {} \;

ENV PATH="$WORKDIR/.local/bin:${PATH}"

RUN chown -R $USERNAME $WORKDIR
RUN chmod 755 $WORKDIR

WORKDIR $WORKDIR
USER $USERNAME