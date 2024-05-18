# Setting Up Tenable Nessus Using Docker

This README provides step-by-step instructions for setting up Tenable Nessus using Docker. You will learn how to fork the repository, clone it to your local machine, pull the Nessus Docker image, and run the container.

## Video Series

### Part 1: 
Video Out Soon


## Prerequisites

Before you begin, ensure you have the following installed:

- Docker: [Install Docker](https://docs.docker.com/get-docker/)
- Git: [Install Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Fork the Repository

1. Go to the [Nessus Docker repository](https://github.com/tenable/nessus-docker) on GitHub.
2. Click on the `Fork` button in the top right corner of the page to create a copy of the repository under your GitHub account.

## Clone the Repository

1. Open your terminal.
2. Clone your forked repository using the following command:

    ```sh
    git clone <repo-url>
    ```

3. Navigate to the cloned repository directory:

    ```sh
    cd nessus-docker
    ```

## Docker

### Pulling the Nessus Docker Image

1. Determine the desired version and OS for the Nessus Docker image. Replace `<version-OS>` in the following command with your specific choice (e.g., `8.15.0-ubuntu`):

    ```sh
    docker pull tenable/nessus:<version-OS>
    ```

### Running the Nessus Docker Container

1. Run the following command to start a new Nessus container. You can customize the ports, volume mappings, and environment variables as needed:

    ```sh
    docker run -d --name nessus -p 8834:8834 tenable/nessus:<version-OS>
    ```

    - `-d`: Run the container in detached mode.
    - `--name nessus`: Assign a name to the container.
    - `-p 8834:8834`: Map port 8834 of the host to port 8834 of the container.

2. (Optional) If you need to persist the data, you can mount a volume to the container:

    ```sh
    docker run -d --name nessus -p 8834:8834 -v /path/to/local/storage:/opt/nessus/var tenable/nessus:<version-OS>
    ```

    Replace `/path/to/local/storage` with the path to the directory on your host where you want to store the data.

## Access Nessus Web Interface

1. Open a web browser and navigate to `https://<your-docker-host-ip>:8834`.
2. Follow the on-screen instructions to complete the Nessus setup.

## Stopping and Removing the Container

To stop the Nessus container, run:

```sh
docker stop nessus
