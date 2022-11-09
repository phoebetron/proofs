# proofs

A collection of completed machine learning tasks using [Getlin]. Once set-up,
simply run the scripts below.

1. Clone this repository in order to get the code on your machine.

```
git clone https://github.com/phoebetron/proofs.git
```

2. Install [Golang], if you have no working environment yet, and download the
   project dependencies.

```
go mod tidy
```

[Getlin]: https://github.com/phoebetron/getlin
[Golang]: https://go.dev



### mnist

Running the proof script with `go run mnist/main.go` produces output similar to
the logs shown below. During the first epochs the Module does yet recognize any
patterns and therefore does not make any predictions. After a while the true
class and the predicted class align more and more often. The script runs for
about two minutes on some MacBook Pro.

```
Downloading mnist_train.csv.zip 100%
Downloaded mnist_train.csv.zip in 1.2s
epo    1        mae 0.181        cla    2        prd    0
epo    2        mae 0.178        cla    0        prd    0
epo    3        mae 0.175        cla    0        prd    0
epo    4        mae 0.175        cla    9        prd    0
epo    5        mae 0.181        cla    6        prd    0
...
epo   96        mae 0.025        cla    3        prd    3
epo   97        mae 0.047        cla    6        prd    6
epo   98        mae 0.031        cla    3        prd    3
epo   99        mae 0.041        cla    4        prd    4
epo  100        mae 0.041        cla    7        prd    7
```

The script prints some example images in the terminal after 100 epochs and shows
what class the image represents and what probabilities the Module produced per
class. The example image below resembles the number `3`.

```







          #####
      ############
     ###############
     #######   ######
     ####       #####
     ##         ####
               #####
             ######
        ##########
       ############
       #############
        ###     ####
                 ###
                 ###
                ####
               ####
        ###########
       ###########
      ###########
       ##   ##



```

The Module predicts a probability of 24.2% for that number to be a `3` and 1.5%
for that number to be an `8`.

```
cla 3
prd [0 0 0 0.24256326 0 0 0 0 0.015090526 0]
```