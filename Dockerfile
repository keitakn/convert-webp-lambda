FROM public.ecr.aws/lambda/nodejs:12

WORKDIR /node/app

RUN npm install -g yarn

COPY package*.json ./

RUN yarn install

COPY . .

CMD ["/bin/bash"]
