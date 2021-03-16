import type { AWS } from '@serverless/typescript';

import convertToWebp from '@functions/convertToWebp';

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
    logRetentionInDays: 3,
    iam: {
      role: {
        statements: [
          {
            Effect: 'Allow',
            Action: ['s3:*'],
            Resource: '*',
          }
        ]
      }
    },
    environment: {
      AWS_NODEJS_CONNECTION_REUSE_ENABLED: '1',
      DEPLOY_STAGE: process.env.DEPLOY_STAGE,
      TRIGGER_BUCKET_NAME: process.env.TRIGGER_BUCKET_NAME,
      DESTINATION_BUCKET_NAME: process.env.DESTINATION_BUCKET_NAME,
    },
    lambdaHashingVersion: '20201221',
  },
  // import the function via paths
  functions: { convertToWebp },
};

module.exports = serverlessConfiguration;
