# StudyBud

StudyBud is a tool for students to better manage their semesters and succeed at their education goals.

## Dev Setup

### Account

Create an account here on GitHub. - GitHub is where the code for StudyBud is kept. We will go over the basics of 'git' later in this introduction. - Basically, this is how we share work between our computers. This allows any developer to make a change to the code and share it with the other people working on the code.

### Computer Setup

1. Install Visual Studio Code on your computer from here: https://code.visualstudio.com/

   - This is the software used to edit the code. This is the recommended code editor for all beginers.
   - Within VS Code, go to the left hand side of the window, and look for icon that looks like four boxes with the one in the top right is floating detached from the others.
   - You will want to install the following extensions (dont worry about what these do. For now, just know you need them):
     - Go
     - Prettier - Code formatter
     - Docker
     - HTML CSS Support
     - MongoDB for VS Code

2. Install Git on your computer from here: https://git-scm.com/downloads

   - This is how we share code to GitHub.
   - Git is a version control system.
   - You will use git to get the latest changes from github, push your changes to github, and manage your branches (dont panic! We will go over this later and its way easier than it seems)

3. Install Golang on your computer from here: https://go.dev/doc/install

   - This is the main language we are using to develop StudyBud. This allows you to develop with Go.

4. Install Docker Desktop on your computer from here: https://www.docker.com/products/docker-desktop/

   - Click 'choose plan' and then 'docker personal'
   - You will need to create an account but then just download and install like normal
   - This is the software we will use to run our code on our laptops.

5. Open a terminal
    - If you are on a Mac, search for 'terminal' in your spotlight search
    - If you are on Windows, search for 'cmd' or 'powershell' in your application
    - Navigate to the folder you want to save the code in
        - tip:
            - on Mac, go to file explorer and navigate to the foler. right click and find services in the list and select 'new terminal at folder'
            - on Windows, go to your file explorer and navigate to the folder. right click and click 'open in terminal'
    - copy and paste this command into your terminal and hit enter to pull down the repo locally
        - git clone https://github.com/WillBallentine/StudyBud.git
    - run this command in your open terminal
        - cd StudyBud
    - leave this terminal open, we will use it later

6. If VS Code is open already, click 'file' and 'open foler' and open the folder you just cloned the repo into in the last step.

After you complete these steps your development environment should be setup and you are ready to code! (if you run into issues, let me know and we can troubleshoot)

### Dev Basics

Ok, at this point, it can feel a bit overwhelming. Don't worry, we are going to take it slow and go item by item. At any point while reading this (or even long after), don't hessitate to ask me any questions. This can feel complicated when its new but trust me, like with anything, the more you do it the more muscle memory you will have.

#### Git

- What is Git? Git is called version control software.
    - This is used to allow multiple people to work on the same code at the same time.
    - Git does not have a user interface but rather works in the terminal.
    - If you are on a Mac, search for 'terminal' in your spotlight search
    - If you are on Windows, search for 'cmd' or 'powershell' in your application


- How do I use git?
    - Open your terminal and type 'git --version'
    - Nice! you just ran a git command!
    - All git commands start with the word 'git' and then have some sort of designator after it. here are a few common ones and what they do:
        - 'git status'
            - this tells you the changes you have locally on your machine.
            - this will also tell you other details about git like what branch you currently have checked out
        - 'git pull'
            - this pulls down the latest version of the code from github for the checked out branch
        - 'git push'
            - this one we will go over more later. basically it will push your changes up to gihub
        - 'git branch -b your_new_branch_name'
            - this creates a new local branch for you to do your work on
            - we NEVER do work on the main branch
        - 'git checkout branch_name'
            - this allows you to switch to the branch that has the name you entered in if it exists
        - 'git add .'
            - this adds all your local changes to your next commit
        - 'git commit -m "your commit message"'
            - this creats a new commit with a custom message. this is what is pushed when you run git push
        - 'git clone link_to_github_repo'

Basic git flow:

1. 'git pull' on the main branch to ensure you have the latest version of the code
2. 'git checkout -b new_branch_name' to create a new branch. names are often related to the work you are doing. for example, if I am working to add a new feature that updates the color of the submit button on the login page, my git command might look something like this 'git checkout -b login_page_button_color_change'
3. do your work in VS Code.
4. when your changes are done go back to the terminal
5. 'git status' to check that your expected changes are there (should show up in red in the terminal)
6. 'git add .' to add all the changes to staged status
7. 'git commit -m "some_message"' to commit your changes and prepare them to be pushed. an example might be 'git commit -m "changed collor of button on login page"'
8. 'git push --set-upstream origin branch_name' so if my branch name was the same as what we setup earlier this would look like this 'git push --set-upstream origin login_page_button_color_change'
9. head to github and create a pull request for others to review and then merge into main


note: never push to main!!!! never work on main!!!