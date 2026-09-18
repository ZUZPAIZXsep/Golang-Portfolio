import { Component, OnInit } from '@angular/core';

import { SkillsService } from '../../services/skills';
import { Skill } from '../../models/skill';

@Component({
  selector: 'app-skills',
  imports: [],
  templateUrl: './skills.html',
  styleUrl: './skills.css'
})
export class Skills implements OnInit {

  skills: Skill[] = [];

  constructor(private skillsService: SkillsService) {}

  ngOnInit(): void {

    this.skillsService.getSkills().subscribe({
      next: (data) => {
        this.skills = data;
      },

      error: (error) => {
        console.error('Failed to load skills:', error);
      }
    });

  }

}