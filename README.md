# DynaTrace account inventory

This project serves as a mapper between accounts and monitored entities. Best practices tell you to plan who will be responsible for onboarding observability for each application. You should plan which team members will deploy the OneAgent, how will you tag monitored entities, and how will you organize your "Management Zones".

If you didn't plan this ahead of time, you may end up with a mess of entities and no way to organize them. This project will help you to organize your entities by mapping them to accounts and teams.

# Idea

The DynaTrace Account API has all the information about accounts, the environment APIs has information about monitored entities, and audit logs have information about who did what and when. This project will use all of these APIs to create a mapping between accounts and monitored entities.

To make visualizing this mapping easier, this project will also use Metabase and MongoDB to create useful queries and dashboards.
