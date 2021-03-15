import type { AWS } from '@serverless/typescript';

import hello from '@functions/hello';

const serverlessConfiguration: AWS = {
  service: 'convert-webp-lambda',
  frameworkVersion: '2',
  custom: {
    webpack: {
      webpackConfig: './webpack.config.js',
      includeModules: true,
    },
    prune: {
      automatic: true,
      number: 1,
    },
  },
  plugins: ['serverless-webpack', 'serverless-prune-plugin'],
  provider: {
    name: 'aws',
    runtime: 'nodejs14.x',
    stage: process.env.DEPLOY_STAGE,
    region: 'ap-northeast-1',
    profile: process.env.DEPLOY_STAGE === 'dev' ? 'nekochans-dev' : 'nekochans-prod',
    iamRoleStatements: [
      {
        Effect: 'Allow',
        Action: ['s3:*'],
        Resource: '*',
      },
    ],
    apiGateway: {
      minimumCompressionSize: 1024,
      shouldStartNameWithService: true,
    },
    environment: {
      AWS_NODEJS_CONNECTION_REUSE_ENABLED: '1',
    },
    lambdaHashingVersion: '20201221',
  },
  // import the function via paths
  functions: { hello },
};

module.exports = serverlessConfiguration;
