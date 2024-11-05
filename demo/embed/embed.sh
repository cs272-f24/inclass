
curl https://api.openai.com/v1/embeddings \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $OPENAI_PROJECT_KEY" \
  -d '{
    "input": "Who teaches CS 272 at USF this semester",
    "model": "text-embedding-3-small"
  }'
