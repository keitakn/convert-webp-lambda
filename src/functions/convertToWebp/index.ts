import { handlerPath } from '@libs/handlerResolver';

export default {
  handler: `${handlerPath(__dirname)}/handler.main`,
  events: [
    {
      s3: {
        bucket: process.env.TRIGGER_BUCKET_NAME,
        event: 's3:ObjectCreated:*',
        rules: [
          { prefix: 'uploads/', suffix: '.png' },
        ],
        existing: true,
      }
    }
  ]
}
