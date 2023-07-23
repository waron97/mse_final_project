# MSE final project

This document contains instructions to run the search engine. Prerequisite for successful local execution is to have Docker installed on the machine. 

Commands are provided for Linux and MacOS systems, but the engine can be run on any OS that supports docker.

After ensuring that Docker is installed and active on the host machine, the first step should be to launch the data collection routine. In practice, this means starting the `db`, `crawler` and `frontier-acceptor` services.

For Linux and MacOS users, this can be achieved with the provided bash script:

```bash
cd src/scripts
./collect-data.sh
```

After collecting at least 1000 documents, the full engine can be started via the following command:

```bash
cd src/scripts
./start-prod.sh
```

In a different terminal, the front-end can also be started. For the front-end, it is necessary to have NodeJS installed and active on the host machine. The instructions for builing and starting the front-end are as follows:

```bash
cd src/web-app
npm install
npm run dev
```

The output from this command will reveal the port on which the front-end is running.

The search engine will not be functional until an index has been build by the `indexer` service. The `indexer` is automatically started alongside the other service when running `start-prod.sh`. When done computing and saving the index, the `indexer` will print a message `"[RunClusteringtask] documents clustered"` the console. At this point, the search engine is fully functional and it should be possible to start inserting queries into the front-end.