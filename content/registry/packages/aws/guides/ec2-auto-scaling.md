---
title: "Auto-Scale EC2 Instances"
meta_desc: "Scale EC2 instances based on demand using AWS Auto Scaling Groups with launch templates and Pulumi."
canonical_url: "https://www.pulumi.com/guides/aws/ec2-auto-scaling"
date: 2025-10-08
category: "Compute"
tags: ["aws", "ec2", "auto-scaling", "high-availability", "elasticity"]
faq:
  - question: What is the difference between launch templates and launch configurations?
    answer: Launch templates are the newer, recommended approach. They support all the latest EC2 features including T2/T3 unlimited mode, spot instances, and more instance types. Launch configurations are deprecated and lack these features. Always use launch templates for new deployments.
  - question: How does Auto Scaling maintain high availability?
    answer: Auto Scaling continuously monitors the health of instances. When an instance fails health checks, Auto Scaling automatically terminates it and launches a replacement to maintain your desired capacity. Distributing instances across multiple availability zones provides additional resilience.
  - question: What is the cooldown period and why does it matter?
    answer: The cooldown period (default 300 seconds) is the time between scaling activities. It prevents Auto Scaling from launching or terminating instances too quickly when metrics fluctuate. Adjust this based on your application's startup time - longer for slow-starting apps.
  - question: How do I add scaling policies to respond to load?
    answer: Use aws.autoscaling.Policy resources to define target tracking or step scaling policies. Target tracking (recommended) automatically adjusts capacity to maintain a metric like CPU utilization at 70%. This example shows the basic group; scaling policies are added separately.
  - question: Can I use Auto Scaling with a load balancer?
    answer: Yes, and it's recommended for production. Attach your Auto Scaling Group to an Application Load Balancer or Network Load Balancer. The load balancer distributes traffic across healthy instances, and you can configure health checks to let Auto Scaling know when to replace unhealthy instances.
---

## How do I auto-scale EC2 instances with AWS Auto Scaling Groups?

**To automatically scale EC2 instances based on demand**, create an Auto Scaling Group with a launch template that defines your instance configuration. Auto Scaling maintains your desired capacity, replaces unhealthy instances, and can scale up or down in response to load.

### Example: Basic Auto Scaling Group

{{< resource-example provider="aws" resource="aws.autoscaling.Group" >}}

## Key configuration details

{{< llm-explain
    resource="aws.autoscaling.Group"
    context="This guide explains how to set up EC2 Auto Scaling Groups with Pulumi"
    prompt="Explain the key configuration parameters for this resource. Return as a markdown list with bold parameter names followed by concise, actionable explanations. Focus on what users need to know for production deployments."
    example="**desiredCapacity**: The number of instances to maintain. Auto Scaling replaces failed instances to maintain this count. Set based on baseline load."
>}}

## Frequently asked questions

**What is the difference between launch templates and launch configurations?**
Launch templates are the newer, recommended approach. They support all the latest EC2 features including T2/T3 unlimited mode, spot instances, and more instance types. Launch configurations are deprecated and lack these features. Always use launch templates for new deployments.

**How does Auto Scaling maintain high availability?**
Auto Scaling continuously monitors the health of instances. When an instance fails health checks, Auto Scaling automatically terminates it and launches a replacement to maintain your desired capacity. Distributing instances across multiple availability zones provides additional resilience.

**What is the cooldown period and why does it matter?**
The cooldown period (default 300 seconds) is the time between scaling activities. It prevents Auto Scaling from launching or terminating instances too quickly when metrics fluctuate. Adjust this based on your application's startup time - longer for slow-starting apps.

**How do I add scaling policies to respond to load?**
Use aws.autoscaling.Policy resources to define target tracking or step scaling policies. Target tracking (recommended) automatically adjusts capacity to maintain a metric like CPU utilization at 70%. This example shows the basic group; scaling policies are added separately.

**Can I use Auto Scaling with a load balancer?**
Yes, and it's recommended for production. Attach your Auto Scaling Group to an Application Load Balancer or Network Load Balancer. The load balancer distributes traffic across healthy instances, and you can configure health checks to let Auto Scaling know when to replace unhealthy instances.
