FROM maven:3-openjdk-11 as buildData

WORKDIR /usr/oracle-blueprint

RUN dpkg --add-architecture i386 && \
 apt-get update;
RUN apt-get install -y libssl-dev:i386
RUN apt-get install git 

RUN git clone --depth 1 https://github.com/PaulsBecks/Blockchain-Logging-Framework

RUN cd Blockchain-Logging-Framework && \
    mvn verify -DskipTests 
    
RUN mkdir output


CMD ["java", "-jar", "Blockchain-Logging-Framework/target/blf-cmd.jar", "validate", "Blockchain-Logging-Framework/src/main/resources/TemplateExample.bcql"]