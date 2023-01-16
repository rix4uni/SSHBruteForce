import random
import string
import argparse
import os
from concurrent.futures import ThreadPoolExecutor

def generate_password():
  # Create a list of possible characters for the password
  characters = list(string.ascii_lowercase + string.ascii_uppercase + string.digits)

  # Shuffle the list of characters to randomize the order
  random.shuffle(characters)

  # Select the first 8 characters from the shuffled list
  password = ''.join(characters[:args.length])

  return password

def generate_passwords(numofpass, num_threads):
  passwords = []
  with ThreadPoolExecutor(max_workers=num_threads) as executor:
    for i in range(numofpass):
      # Generate a password asynchronously using a thread
      password_future = executor.submit(generate_password)
      passwords.append(password_future)

  # Wait for all threads to complete and return the passwords
  return [password_future.result() for password_future in passwords]

if __name__ == '__main__':
  # Use argparse to parse command-line arguments
  parser = argparse.ArgumentParser(description='Generate random passwords')
  parser.add_argument('-l', '--length', type=int, default=8, help='Number of passwords length to generate (default 8)')
  parser.add_argument('-n', '--numofpass', type=int, default=1000, help='Number of passwords to generate (default 1000)')
  parser.add_argument('-t','--threads', type=int, default=100, help='Number of threads to use  (default 100)')
  parser.add_argument('-o','--output', default='ssh-password.txt', help='Output file for passwords (default ssh-password.txt)')
  args = parser.parse_args()

  # Generate the specified number of passwords concurrently
  passwords = generate_passwords(args.numofpass, args.threads)

  # If an output file is specified, write the passwords to the file
  if args.output:
    with open(args.output, 'w') as f:
      for password in passwords:
        f.write(password + '\n')
  else:
    # Otherwise, print the passwords to the console
    for password in passwords:
      print(password)