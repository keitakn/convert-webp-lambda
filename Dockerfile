FROM public.ecr.aws/lambda/nodejs:12

WORKDIR /node/app

COPY package*.json ./

RUN npm install

COPY . .

CMD ["/bin/bash"]
