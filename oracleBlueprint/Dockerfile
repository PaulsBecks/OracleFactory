FROM maven:3-openjdk-11 as buildData

WORKDIR /usr/oracle-blueprint

RUN apt-get install git

RUN git clone --depth 1 https://github.com/TU-ADSP/Blockchain-Logging-Framework

RUN cd Blockchain-Logging-Framework && \
    mvn verify -DskipTests 
    
RUN mkdir output


CMD ["java", "-jar", "Blockchain-Logging-Framework/target/blf-cmd.jar", "extract", "Blockchain-Logging-Framework/src/main/resources/TemplateExample.bcql"]