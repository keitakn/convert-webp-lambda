import 'source-map-support/register';

import type { S3CreateEvent, S3Handler } from "aws-lambda"

import { middyfy } from '@libs/lambda';

const convertToWebp: S3Handler = async (event: S3CreateEvent, context, callback) => {

  console.log('🐱');
  console.log(event);
  console.log(context);
  console.log('🐱');

  return callback();
}

export const main = middyfy(convertToWebp);
