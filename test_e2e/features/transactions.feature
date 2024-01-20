Feature: Transaction processing
    In order to inform my clients
    As Stori platform
    I need to process a transaction file
    And send an email to my clients with the results

  Scenario: Process local CSV file
    Given there is a local CSV file with following data
      | Id | Date | Transaction |
      | 0  | 7/15 | +60.5       |
      | 1  | 7/28 | -10.3       |
      | 2  | 8/2  | -20.46      |
      | 3  | 8/13 | +10         |
    When the system is executed
    Then I receive an email with subject "Stori transaction summary" and with the following information
      | Total balance is: 39.74             |
      | Number of transactions in July: 2   |
      | Number of transactions in August: 2 |
      | Average debit amount: -15.38        |
      | Average credit amount: 35.25        |

  Scenario: Process S3 CSV file
    Given there is a S3 CSV file with following data
      | Id | Date | Transaction |
      | 0  | 7/15 | +60.5       |
      | 1  | 7/28 | -10.3       |
      | 2  | 8/2  | -20.46      |
      | 3  | 8/13 | +10         |
    When the system is executed
    Then I receive an email with subject "Stori transaction summary" and with the following information
      | Total balance is: 39.74             |
      | Number of transactions in July: 2   |
      | Number of transactions in August: 2 |
      | Average debit amount: -15.38        |
      | Average credit amount: 35.25        |