# RPSWEChallenge (RE Partners Software Engineering Challenge)

## Reproduction

1. [https://docs.amplify.aws/gen1/react/start/getting-started/installation/](https://docs.amplify.aws/gen1/react/start/getting-started/installation/)
   1. npm install -g @aws-amplify/cli
   2. amplify configure
      1. Use eu-central-1 region
      2. Instead of the AdministratorAccess-Amplify policy that is recommended to be attached to the user, use the AdministratorAccess policy because the former doesn’t have all permissions for us to deploy a serverless container later
2. [https://docs.amplify.aws/gen1/react/start/getting-started/setup/](https://docs.amplify.aws/gen1/react/start/getting-started/setup/)
   1. amplify init
      1. Do you want to continue with Amplify Gen 1? y
   2. npm create vite@latest
      1. Project name: rpswechallenge
      2. Select a framework: React
      3. Select a variant: TypeScript
3. cd rpswechallenge
4. [https://docs.amplify.aws/gen1/react/start/getting-started/hosting/](https://docs.amplify.aws/gen1/react/start/getting-started/hosting/)
5. [https://docs.amplify.aws/gen1/react/tools/cli/usage/containers/](https://docs.amplify.aws/gen1/react/tools/cli/usage/containers/)
   1. amplify configure project
      1. Select “Advanced: Container-based deployments”
      2. Do you want to enable container-based deployments? y
   2. Skip “amplify add storage”
   3. amplify add api
      1. REST
      2. API Gateway + AWS Fargate (Container=based)
      3. orderpackscalculatorapi
      4. Custom (bring your own Dockerfile or docker-compose.yml)
      5. On every “amplify push” (Fully managed container source)
      6. Do you want to restrict API access? n
   4. Download Go 1.24.3 from [https://go.dev/dl/](https://go.dev/dl/)
   5. [https://gin-gonic.com/en/docs/quickstart/](https://gin-gonic.com/en/docs/quickstart/)
      1. “go mod init orderpackscalculatorapi” in amplify/backend/orderpackscalculatorapi/src instead of “go mod init gin”
      2. Skip steps 2-4
   6. [https://hub.docker.com/\_/golang](https://hub.docker.com/_/golang)

## Local development and testing

[https://docs.amplify.aws/gen1/react/tools/cli/usage/containers/#local-development-and-testing](https://docs.amplify.aws/gen1/react/tools/cli/usage/containers/#local-development-and-testing)
