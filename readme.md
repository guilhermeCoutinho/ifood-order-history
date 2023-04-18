<a href="https://hub.docker.com/repository/docker/guilhermecoutinho/ifood-order-history"><img alt="Docker Cloud Build Status" src="https://img.shields.io/docker/cloud/build/guilhermecoutinho/ifood-order-history" ></a>
# Ifood order history exporter
This program exports ifood's order history to a tabulated format separated by tabs. 
You can [customize which data you want to expot here](https://github.com/guilhermeCoutinho/ifood-order-history/blob/master/utils.go#L21). The default behaviour looks like this

| Date of order  | Establishment name | Price | Orderem items separated by a semi-colon
| ------------- | ------------- | ------------- | ------------- |
| 1-August-2020 |	Ice cream store |	29.97 |	0001 - 1 ice cream cone (2 flavours); 0002 -  1 ice cream cone (3 flavours);

## Usage

### Authenticating
You will need to authenticate with a browser first because of Okta. Maybe I figure a way to trigger the ifood auth popup from the terminal some day. So the idea is to copy the authentication token request from the developer options and run it. I choose to do it with curl. The image should illustrate this. 

After that, [paste the authentication token here](https://github.com/guilhermeCoutinho/ifood-order-history/blob/master/constants.go#L8)

![Screen Shot 2021-04-07 at 19 45 09](https://user-images.githubusercontent.com/7122366/113944937-c09bcc80-97db-11eb-9d8b-920886194b76.png)

---

### Running the program

#### Clone and run
Dependencies: golang
To run it, first [paste the authentication token here](https://github.com/guilhermeCoutinho/ifood-order-history/blob/master/constants.go#L8)
Then execute:

```make run```

#### Using docker 
<a href="https://hub.docker.com/r/guilhermecoutinho/ifood-order-history/tags">![Docker Image Version (latest semver)](https://img.shields.io/docker/v/guilhermecoutinho/ifood-order-history?color=orange&label=container)</a>

The recent changes made to the authentication method broke the docker build. Updates will come soon to fix this.


### Changing the filter
I suggest you take a look at the *.go files* to see what you can do with it. [Change the filter here](https://github.com/guilhermeCoutinho/ifood-order-history/blob/master/main.go#L32):
	
```
filter := func(o *Order) bool {
    return o.LastStatus == "CONCLUDED" &&
        o.CreatedAt.After(startingDate) &&
        o.CreatedAt.Before(endDate)
}
```
