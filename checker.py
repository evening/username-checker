#!/usr/bin/python

import requests
import sys
from urllib.parse import urlparse
from colorama import Fore, Style, init
import multiprocessing

init(convert=True)

def check(url, word, output):
    r = requests.get(url.format(
        word)) if "{}" in url else requests.get(url + word)
    if r.status_code == 404:
        with open(output, 'a') as f:
            f.write(word + "\n")
        print(f"{Fore.GREEN}[+]{Style.RESET_ALL} {word} is available")
    else:
        print(f"{Fore.RED}[-]{Style.RESET_ALL} {word} is not available")


def main(url, dictionary, output):
    p = urlparse(url)
    url = ''.join(['http://', p.netloc, p.path])
    dict = open(dictionary, 'r')
    pool = multiprocessing.Pool()
    for word in dict:
        word = word.rstrip()
        pool.apply_async(check, args=(url, word, output))
    dict.close()
    pool.close()
    pool.join()

if __name__ == '__main__':
    if len(sys.argv) == 4:
        main(sys.argv[1], sys.argv[2], sys.argv[3])
    else:
        print("Usage:   checker.py [url] [input] [output]\n"
              "Example: checker.py github.com/{} dict github_example"
              )
